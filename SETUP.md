# GitHub Visitor Counter - Developer Setup

This guide is for developers who want to set up and run the GitHub Visitor Counter project locally or deploy their own instance.

## 🛠️ Tech Stack

- **Backend**: Go with Fiber web framework
- **Database**: Supabase (PostgreSQL)
- **Environment**: Docker-ready
- **Deployment**: Vercel, Railway, Heroku compatible

## 📋 Prerequisites

- Go 1.19 or higher
- Supabase account
- Git
- (Optional) Docker for containerized deployment

## 🚀 Local Development Setup

### 1. Clone the Repository

```bash
git clone https://github.com/MohammedAbidNafi/github-visitor-counter.git
cd github-visitor-counter
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Environment Configuration

Create a `.env` file in the root directory:

```env
SUPABASE_URL="your_supabase_project_url"
SUPABASE_KEY="your_supabase_anon_key"
```

**Getting Supabase Credentials:**

1. Go to [supabase.com](https://supabase.com)
2. Create a new project
3. Go to Settings → API
4. Copy the Project URL and anon/public key

### 4. Database Setup

Connect to your Supabase project and run this SQL:

```sql
-- Create the visits table
CREATE TABLE visits (
  id TEXT NOT NULL,
  fingerprint TEXT NOT NULL,
  visited_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Add indexes for better performance
CREATE INDEX idx_visits_id ON visits(id);
CREATE INDEX idx_visits_visited_at ON visits(visited_at);
CREATE INDEX idx_visits_fingerprint ON visits(fingerprint);

-- Optional: Add RLS (Row Level Security) if needed
-- ALTER TABLE visits ENABLE ROW LEVEL SECURITY;
```

### 5. Run the Application

```bash
# Development mode
go run .

# Or build and run
go build -o visitor-counter
./visitor-counter
```

The server will start on `http://localhost:3000`

### 6. Test the API

```bash
# Test basic endpoint
curl http://localhost:3000/testuser

# Test with parameters
curl "http://localhost:3000/testuser?label=Test&bg_color=FF5733&rounded=true"
```

## 📁 Project Structure

```
github-visitor-counter/
├── main.go              # Main application file
├── go.mod              # Go modules file
├── go.sum              # Go modules checksum
├── .env                # Environment variables (create this)
├── README.md           # User documentation
├── SETUP.md            # This developer guide
└── .gitignore          # Git ignore file
```

## 🔧 Configuration Options

### Environment Variables

| Variable       | Required | Description                   |
| -------------- | -------- | ----------------------------- |
| `SUPABASE_URL` | Yes      | Your Supabase project URL     |
| `SUPABASE_KEY` | Yes      | Your Supabase anon/public key |
| `PORT`         | No       | Server port (default: 3000)   |

### Application Settings

You can modify these in `main.go`:

- **Default digits**: Change the default from 6 to any number
- **Rate limiting**: Add rate limiting middleware
- **CORS settings**: Configure allowed origins
- **Caching**: Add Redis or in-memory caching

## 🚀 Deployment Options

### Deploy to Vercel

1. **Push to GitHub**:

```bash
git add .
git commit -m "Initial commit"
git push origin main
```

2. **Deploy on Vercel**:

   - Go to [vercel.com](https://vercel.com)
   - Import your GitHub repository
   - Add environment variables in dashboard
   - Deploy

3. **Vercel Configuration** (optional `vercel.json`):

```json
{
  "functions": {
    "main.go": {
      "runtime": "vercel-go@3.0.0"
    }
  }
}
```

### Deploy to Railway

1. **Connect Repository**:

   - Go to [railway.app](https://railway.app)
   - Connect your GitHub repository

2. **Add Environment Variables**:

   - Set `SUPABASE_URL`
   - Set `SUPABASE_KEY`

3. **Deploy**: Railway will automatically detect Go and deploy

### Deploy to Heroku

1. **Create Heroku App**:

```bash
heroku create your-app-name
```

2. **Set Environment Variables**:

```bash
heroku config:set SUPABASE_URL="your_url"
heroku config:set SUPABASE_KEY="your_key"
```

3. **Deploy**:

```bash
git push heroku main
```

### Docker Deployment

Create a `Dockerfile`:

```dockerfile
FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 3000
CMD ["./main"]
```

Build and run:

```bash
docker build -t visitor-counter .
docker run -p 3000:3000 --env-file .env visitor-counter
```

## 🧪 Testing

### Manual Testing

```bash
# Basic functionality
curl http://localhost:3000/testuser

# With repository
curl "http://localhost:3000/testuser?repo=test-repo"

# Full customization
curl "http://localhost:3000/testuser?label=Views&digits=4&bg_color=FF5733&text_color=FFFFFF&rounded=true&unique=true&days=30"
```

### Load Testing (Optional)

Use tools like `ab` or `wrk`:

```bash
# Apache Bench
ab -n 1000 -c 10 http://localhost:3000/testuser

# wrk
wrk -t12 -c400 -d30s http://localhost:3000/testuser
```

## 🔍 Monitoring and Logs

### Application Logs

The application logs:

- Supabase connection status
- Insert errors
- Parse errors
- Server start/stop

### Database Monitoring

Monitor your Supabase dashboard for:

- Query performance
- Storage usage
- Connection counts

## 🛡️ Security Considerations

1. **Environment Variables**: Never commit `.env` files
2. **Rate Limiting**: Consider adding rate limiting for production
3. **CORS**: Configure appropriate CORS settings
4. **Database**: Use RLS (Row Level Security) if needed
5. **Monitoring**: Set up error tracking (Sentry, etc.)

## 🔧 Customization

### Adding New Parameters

1. **Add parameter parsing** in `trackAndRender()`:

```go
newParam := c.Query("new_param", "default_value")
```

2. **Pass to SVG builder**:

```go
svg := buildSVG(label, count, bgColor, textColor, rounded, digits, newParam)
```

3. **Update `buildSVG()` function** to handle the new parameter

### Modifying SVG Output

Edit the `buildSVG()` function to:

- Change dimensions
- Add animations
- Modify styling
- Add new visual elements

## 📊 Database Schema Details

```sql
-- The visits table structure
CREATE TABLE visits (
  id TEXT NOT NULL,                    -- username or username+repo
  fingerprint TEXT NOT NULL,           -- SHA256 hash of IP + User-Agent
  visited_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for performance
CREATE INDEX idx_visits_id ON visits(id);
CREATE INDEX idx_visits_visited_at ON visits(visited_at);
CREATE INDEX idx_visits_fingerprint ON visits(fingerprint);
```

**Field Explanations:**

- `id`: Identifier for tracking (username or username+repo)
- `fingerprint`: Anonymous visitor identification
- `visited_at`: Timestamp for time-based filtering

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes
4. Test thoroughly
5. Commit: `git commit -am 'Add new feature'`
6. Push: `git push origin feature-name`
7. Submit a Pull Request

## 📝 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🆘 Troubleshooting

### Common Issues

**Supabase Connection Failed**:

- Check your `.env` file
- Verify Supabase credentials
- Ensure table exists

**SVG Not Displaying**:

- Check Content-Type header
- Verify SVG syntax
- Test URL directly in browser

**High Response Times**:

- Add database indexes
- Implement caching
- Consider connection pooling

### Debug Mode

Add debug logging:

```go
log.Println("Debug: Processing request for", username)
```

---

Need help? Open an issue on [GitHub](https://github.com/MohammedAbidNafi/github-visitor-counter/issues)!
