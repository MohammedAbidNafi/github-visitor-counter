# GitHub Visitor Counter

A free service to track and display visitor counts on your GitHub profile and repositories with beautiful, customizable SVG badges.

## 🚀 Quick Start

Simply add this to your GitHub README:

```markdown
![Profile Views](https://viewcounter.live/your-github-username)
```

Replace `your-github-username` with your actual GitHub username and you're done!

## ✨ Features

- 🔢 **6-digit display** with leading zeros by default
- 📊 **Unique visitor tracking**
- 📅 **Time-based filtering** (last N days)
- 🎨 **Fully customizable** colors and styling
- 📍 **Repository-specific** counters
- �️ **SVG format** - works everywhere
- 🆓 **Completely free** to use

## 📖 Usage Examples

### Basic Profile Counter

```markdown
![Profile Views](https://viewcounter.live/johndoe)
```

### Repository-Specific Counter

```markdown
![Repo Views](https://viewcounter.live/johndoe?repo=my-awesome-project)
```

### With Custom Label

```markdown
![Profile Views](https://viewcounter.live/johndoe?label=Profile%20Views)
```

### Styled Counter

```markdown
![Profile Views](https://viewcounter.live/johndoe?label=Visitors&bg_color=0366d6&text_color=FFFFFF&rounded=true)
```

## 🎨 Customization Options

| Parameter    | Description                  | Example                  | Default  |
| ------------ | ---------------------------- | ------------------------ | -------- |
| `label`      | Text displayed above counter | `?label=Profile%20Views` | None     |
| `digits`     | Number of digits to show     | `?digits=4`              | 6        |
| `bg_color`   | Background color (hex)       | `?bg_color=FF5733`       | 000000   |
| `text_color` | Text color (hex)             | `?text_color=FFFFFF`     | FFFFFF   |
| `rounded`    | Rounded corners              | `?rounded=true`          | false    |
| `unique`     | Count unique visitors only   | `?unique=true`           | false    |
| `days`       | Show last N days only        | `?days=30`               | All time |
| `repo`       | Track specific repository    | `?repo=project-name`     | Profile  |

## 🎯 Real Examples

### GitHub Profile

```markdown
![Profile Views](https://viewcounter.live/MohammedAbidNafi?label=Profile%20Views&bg_color=0366d6&text_color=FFFFFF&rounded=true)
```

### Project Repository

```markdown
![Repo Views](https://viewcounter.live/MohammedAbidNafi?repo=github-visitor-counter&label=Repository%20Views&bg_color=28a745&text_color=FFFFFF&rounded=true&unique=true)
```

### Minimalist Style

```markdown
![Views](https://viewcounter.live/MohammedAbidNafi?digits=4&bg_color=000000&text_color=FFFFFF)
```

## 🌈 Popular Color Schemes

### GitHub Blue

```
?bg_color=0366d6&text_color=FFFFFF
```

### Success Green

```
?bg_color=28a745&text_color=FFFFFF
```

### Dark Theme

```
?bg_color=24292e&text_color=FFFFFF
```

### Warning Orange

```
?bg_color=fd7e14&text_color=FFFFFF
```

### Purple

```
?bg_color=6f42c1&text_color=FFFFFF
```

## 📊 Display Format

The counter shows as individual digit blocks in a clean, vertical layout:

- **Label** (optional): Appears at the top
- **Counter**: Individual digit blocks below
- **Default**: 6 digits with leading zeros (e.g., `000042`)

## 🔗 Combining Parameters

You can combine multiple parameters using `&`:

```markdown
![Custom Counter](https://viewcounter.live/username?label=Total%20Views&repo=project&digits=5&bg_color=FF6B6B&text_color=FFFFFF&rounded=true&unique=true&days=365)
```

## 💡 Tips

1. **URL Encode Spaces**: Use `%20` for spaces in labels
2. **Hex Colors**: Don't include the `#` symbol
3. **Repository Names**: Use exact repository name (case-sensitive)
4. **Unique Visitors**: Use `unique=true` for more accurate engagement metrics
5. **Time Filtering**: Use `days` parameter for recent activity tracking

## 🤝 Support

- Free to use for all GitHub users
- No registration required
- No rate limits
- Works in all README files, wikis, and markdown documents

## 🔥 Trending Examples

**Animated Style (using different colors):**

```markdown
![Day Views](https://viewcounter.live/username?days=1&label=Today&bg_color=FF6B6B)
![Week Views](https://viewcounter.live/username?days=7&label=This%20Week&bg_color=4ECDC4)
![Month Views](https://viewcounter.live/username?days=30&label=This%20Month&bg_color=45B7D1)
```

**Repository Showcase:**

```markdown
## 📊 Repository Stats

![Total Views](https://viewcounter.live/username?repo=project&label=Total%20Views&bg_color=2C3E50&text_color=FFFFFF&rounded=true)
![Unique Visitors](https://viewcounter.live/username?repo=project&label=Unique%20Visitors&unique=true&bg_color=E74C3C&text_color=FFFFFF&rounded=true)
```

---

Start tracking your GitHub visitors now! Just replace the username in any example above and add it to your README. 🎉
