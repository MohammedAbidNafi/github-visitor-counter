# GitHub Visitor Counter

A free service to track and display visitor counts on your GitHub profile and repositories with beautiful, customizable SVG badges.

## 🚀 Quick Start

Simply add this to your Git## 🙏 Acknowledgments

- Built with ❤️ using [Go](https://golang.org/) and [Fiber](https://gofiber.io/) web framework
- Database powered by [Supabase](https://supabase.com/) - The open source Firebase alternative
- Deployed and hosted on [Railway](https://railway.app/) - Modern application deployment platform
- Inspired by the community need for simple, privacy-focused visitor tracking
- Thanks to all contributors and users who help improve this service
- Special thanks to the open source community for making tools like this possible

## 💖 Support This Project

If you find this service useful, consider supporting its development and hosting costs:

[![Sponsor](https://img.shields.io/badge/Sponsor-❤️-ff69b4?style=for-the-badge&logo=github-sponsors)](https://github.com/sponsors/MohammedAbidNafi)

Your support helps keep this service free for everyone! 🙏DME:

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

![Profile Views](https://viewcounter.live/johndoe)

### Repository-Specific Counter

```markdown
![Repo Views](https://viewcounter.live/johndoe?repo=my-awesome-project)
```

![Repo Views](https://viewcounter.live/johndoe?repo=my-awesome-project)

### With Custom Label

```markdown
![Profile Views](https://viewcounter.live/johndoe?label=Profile%20Views)
```

![Profile Views](https://viewcounter.live/johndoe?label=Profile%20Views)

### Styled Counter

```markdown
![Profile Views](https://viewcounter.live/johndoe?label=Visitors&bg_color=0366d6&text_color=FFFFFF&rounded=true)
```

![Profile Views](https://viewcounter.live/johndoe?label=Visitors&bg_color=0366d6&text_color=FFFFFF&rounded=true)

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

![Profile Views](https://viewcounter.live/MohammedAbidNafi?label=Profile%20Views&bg_color=0366d6&text_color=FFFFFF&rounded=true)

### Project Repository

```markdown
![Repo Views](https://viewcounter.live/MohammedAbidNafi?repo=github-visitor-counter&label=Repository%20Views&bg_color=28a745&text_color=FFFFFF&rounded=true&unique=true)
```

![Repo Views](https://viewcounter.live/MohammedAbidNafi?repo=github-visitor-counter&label=Repository%20Views&bg_color=28a745&text_color=FFFFFF&rounded=true&unique=true)

### Minimalist Style

```markdown
![Views](https://viewcounter.live/MohammedAbidNafi?digits=4&bg_color=000000&text_color=FFFFFF)
```

![Views](https://viewcounter.live/MohammedAbidNafi?digits=4&bg_color=000000&text_color=FFFFFF)

## 🌈 Popular Color Schemes

### GitHub Blue

```
?bg_color=0366d6&text_color=FFFFFF
```

![GitHub Blue](https://viewcounter.live/demo?label=GitHub%20Blue&bg_color=0366d6&text_color=FFFFFF&rounded=true)

### Success Green

```
?bg_color=28a745&text_color=FFFFFF
```

![Success Green](https://viewcounter.live/demo?label=Success%20Green&bg_color=28a745&text_color=FFFFFF&rounded=true)

### Dark Theme

```
?bg_color=24292e&text_color=FFFFFF
```

![Dark Theme](https://viewcounter.live/demo?label=Dark%20Theme&bg_color=24292e&text_color=FFFFFF&rounded=true)

### Warning Orange

```
?bg_color=fd7e14&text_color=FFFFFF
```

![Warning Orange](https://viewcounter.live/demo?label=Warning%20Orange&bg_color=fd7e14&text_color=FFFFFF&rounded=true)

### Purple

```
?bg_color=6f42c1&text_color=FFFFFF
```

![Purple](https://viewcounter.live/demo?label=Purple&bg_color=6f42c1&text_color=FFFFFF&rounded=true)

## 📊 Display Format

The counter shows as individual digit blocks in a clean, vertical layout:

- **Label** (optional): Appears at the top
- **Counter**: Individual digit blocks below
- **Default**: 6 digits with leading zeros (e.g., `000042`)

### Digit Variations

**4 Digits:**
![4 Digits](https://viewcounter.live/demo?digits=4&label=4%20Digits&bg_color=333333&text_color=FFFFFF&rounded=true)

**6 Digits (Default):**
![6 Digits](https://viewcounter.live/demo?digits=6&label=6%20Digits&bg_color=0366d6&text_color=FFFFFF&rounded=true)

**8 Digits:**
![8 Digits](https://viewcounter.live/demo?digits=8&label=8%20Digits&bg_color=28a745&text_color=FFFFFF&rounded=true)

## 🔗 Combining Parameters

You can combine multiple parameters using `&`:

```markdown
![Custom Counter](https://viewcounter.live/username?label=Total%20Views&repo=project&digits=5&bg_color=FF6B6B&text_color=FFFFFF&rounded=true&unique=true&days=365)
```

![Custom Counter](https://viewcounter.live/demo?label=Total%20Views&digits=5&bg_color=FF6B6B&text_color=FFFFFF&rounded=true)

## 💡 Tips

1. **URL Encode Spaces**: Use `%20` for spaces in labels
2. **Hex Colors**: Don't include the `#` symbol
3. **Repository Names**: Use exact repository name (case-sensitive)
4. **Unique Visitors**: Use `unique=true` for more accurate engagement metrics
5. **Time Filtering**: Use `days` parameter for recent activity tracking

## 🔥 Trending Examples

**Time-based Tracking (using different colors):**

```markdown
![Day Views](https://viewcounter.live/username?days=1&label=Today&bg_color=FF6B6B&text_color=FFFFFF&rounded=true)
![Week Views](https://viewcounter.live/username?days=7&label=This%20Week&bg_color=4ECDC4&text_color=FFFFFF&rounded=true)
![Month Views](https://viewcounter.live/username?days=30&label=This%20Month&bg_color=45B7D1&text_color=FFFFFF&rounded=true)
```

![Day Views](https://viewcounter.live/demo?label=Today&bg_color=FF6B6B&text_color=FFFFFF&rounded=true&digits=3)
![Week Views](https://viewcounter.live/demo?label=This%20Week&bg_color=4ECDC4&text_color=FFFFFF&rounded=true&digits=4)
![Month Views](https://viewcounter.live/demo?label=This%20Month&bg_color=45B7D1&text_color=FFFFFF&rounded=true&digits=4)

**Repository Analytics Dashboard:**

```markdown
## 📊 Repository Stats

![Total Views](https://viewcounter.live/username?repo=project&label=Total%20Views&bg_color=2C3E50&text_color=FFFFFF&rounded=true)
![Unique Visitors](https://viewcounter.live/username?repo=project&label=Unique%20Visitors&unique=true&bg_color=E74C3C&text_color=FFFFFF&rounded=true)
![Recent Activity](https://viewcounter.live/username?repo=project&label=Last%2030%20Days&days=30&bg_color=8E44AD&text_color=FFFFFF&rounded=true)
```

![Total Views](https://viewcounter.live/demo?label=Total%20Views&bg_color=2C3E50&text_color=FFFFFF&rounded=true)
![Unique Visitors](https://viewcounter.live/demo?label=Unique%20Visitors&bg_color=E74C3C&text_color=FFFFFF&rounded=true)
![Recent Activity](https://viewcounter.live/demo?label=Last%2030%20Days&bg_color=8E44AD&text_color=FFFFFF&rounded=true)

## 🤝 Support

- Free to use for all GitHub users
- No registration required
- No rate limits
- Works in all README files, wikis, and markdown documents

## � Acknowledgments

- Built with ❤️ using [Go](https://golang.org/) and [Fiber](https://gofiber.io/) web framework
- Database powered by [Supabase](https://supabase.com/) - The open source Firebase alternative
- Inspired by the community need for simple, privacy-focused visitor tracking
- Thanks to all contributors and users who help improve this service
- Special thanks to the open source community for making tools like this possible

## 📄 License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## � Links

- 🌐 **Live Service**: [viewcounter.live](https://viewcounter.live)
- 📂 **Source Code**: [GitHub Repository](https://github.com/MohammedAbidNafi/github-visitor-counter)
- 🐛 **Report Issues**: [GitHub Issues](https://github.com/MohammedAbidNafi/github-visitor-counter/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/MohammedAbidNafi/github-visitor-counter/discussions)

---

Start tracking your GitHub visitors now! Just replace the username in any example above and add it to your README. 🎉
