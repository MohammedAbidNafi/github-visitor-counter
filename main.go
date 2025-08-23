package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/supabase-go"
)

var client *supabase.Client

func main() {

	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatalf("Error loading .env file: %v", err_env)
	}

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")
	print("Supabase url"+url, "Supabase key: "+key)

	var err error
	client, err = supabase.NewClient(url, key, &supabase.ClientOptions{})
	if err != nil {
		log.Fatal("Supabase client error:", err)

	}

	app := fiber.New()

	app.Get("/:username", trackAndRender)

	log.Fatal(app.Listen(":3000"))
}

func trackAndRender(c *fiber.Ctx) error {
	username := c.Params("username")
	repo := c.Query("repo", "")
	id := username
	if repo != "" {
		id = username + "+" + repo
	}

	fingerprint := hashFingerprint(c)
	_, _, err := client.From("visits").
		Insert(map[string]interface{}{
			"id":          id,
			"fingerprint": fingerprint,
		}, false, "", "", "representation").
		Execute()
	if err != nil {
		log.Println("Insert error:", err)
	}

	// params
	days := c.QueryInt("days", 0)
	unique := c.QueryBool("unique", false)
	label := c.Query("label", "")
	bgColor := "#" + c.Query("bg_color", "000000")
	textColor := "#" + c.Query("text_color", "FFFFFF")
	rounded := c.QueryBool("rounded", false)
	digits := c.QueryInt("digits", 6)

	data, _, err := client.From("visits").Select("fingerprint, visited_at", "exact", false).Eq("id", id).Execute()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("DB error")
	}

	type Visit struct {
		Fingerprint string    `json:"fingerprint"`
		VisitedAt   time.Time `json:"visited_at"`
	}
	var visits []Visit

	if days > 0 {

		since := time.Now().AddDate(0, 0, -days)
		data, _, err = client.From("visits").Select("fingerprint, visited_at", "exact", false).
			Eq("id", id).
			Gte("visited_at", since.Format(time.RFC3339)).
			Execute()
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("DB error")
		}
	}

	err = json.Unmarshal(data, &visits)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Parse error")
	}

	count := 0
	if unique {
		uniqSet := map[string]bool{}
		for _, v := range visits {
			uniqSet[v.Fingerprint] = true
		}
		count = len(uniqSet)
	} else {
		count = len(visits)
	}

	svg := buildSVG(label, count, bgColor, textColor, rounded, digits)
	c.Set("Content-Type", "image/svg+xml")
	return c.SendString(svg)
}

func hashFingerprint(c *fiber.Ctx) string {
	data := c.IP() + c.Get("User-Agent")
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func buildSVG(label string, count int, bgColor, textColor string, rounded bool, digits int) string {

	countStr := fmt.Sprintf("%0*d", digits, count)
	blockRadius := "0"
	if rounded {
		blockRadius = "8"
	}

	// spacing
	blockWidth := 40
	blockHeight := 50
	padding := 5

	totalWidth := len(countStr)*(blockWidth+padding) + padding
	labelHeight := 0
	totalHeight := blockHeight + 2*padding

	if label != "" {
		labelHeight = 30
		totalHeight += labelHeight + padding
	}

	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d">`,
		totalWidth, totalHeight)

	// label
	if label != "" {
		svg += fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" fill="%s" font-size="16" font-family="Arial" dominant-baseline="middle">%s</text>`,
			totalWidth/2, labelHeight/2+padding, textColor, label)
	}

	startY := padding + labelHeight
	if label != "" {
		startY += padding
	}

	for i, ch := range countStr {
		x := padding + i*(blockWidth+padding)
		svg += fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" rx="%s" ry="%s" fill="%s"/>`,
			x, startY, blockWidth, blockHeight, blockRadius, blockRadius, bgColor)

		svg += fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" fill="%s" font-size="20" font-family="Arial" dominant-baseline="middle">%c</text>`,
			x+blockWidth/2, startY+blockHeight/2, textColor, ch)
	}

	svg += `</svg>`
	return svg
}
