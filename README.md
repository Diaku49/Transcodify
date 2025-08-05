# 🎥 VideoTranscodeApp

Hey there! 👋 This is my video transcoding app - basically a cool way to upload videos and get them converted into different formats. Think of it as your personal video processing buddy!

## What's this all about? 🤔

So you've got a video file that needs to be in a different format? Maybe you want to convert that huge 4K video to something more manageable, or you need it in a specific format for your project? That's exactly what this app does!

Upload your video, pick your desired format, and let the magic happen. The app handles all the heavy lifting in the background while you grab a coffee.

## 🏗️ How it's built

This is a full-stack app with some pretty cool tech behind it:

### Frontend
- **React 19** - Because who doesn't love React?
- **React Router** - For smooth navigation
- **Axios** - Making those API calls
- **React Toastify** - For those nice little notifications
- **CSS Modules** - Keeping styles organized and clean

### Backend
- **Go** - Fast, efficient, and reliable
- **Chi Router** - For handling all those HTTP requests
- **GORM** - Database stuff made easy
- **PostgreSQL** - Rock-solid database
- **JWT** - Keeping things secure
- **Redis** - For caching and session management
- **RabbitMQ** - Message queuing for background tasks

### Worker Service
- **Go** - Processing videos in the background
- **FFmpeg** - The real MVP for video transcoding
- **AWS S3** - Storing all those video files
- **RabbitMQ** - Coordinating the work

## 🚀 What can you do?

- **Upload videos** - Drag, drop, and go!
- **Convert formats** - MP4, AVI, MOV, you name it
- **User accounts** - Keep track of your uploads
- **Background processing** - No waiting around
- **Email notifications** - Know when your video is ready

## 🛠️ Getting it running

### Prerequisites
- Go 1.24+
- Node.js & npm
- PostgreSQL
- Redis
- RabbitMQ
- FFmpeg

### Quick start
1. Clone the repo
2. Set up your environment variables
3. Run the backend: `cd backend && go run cmd/main.go`
4. Run the worker: `cd worker && go run main.go`
5. Run the frontend: `cd frontend && npm start`

## The architecture

```
Frontend (React) → Backend (Go) → Worker (Go)
                    ↓
                RabbitMQ → Redis → PostgreSQL
                    ↓
                AWS S3 (storage)
```

## License

This is just a personal project, so feel free to use it as inspiration for your own stuff!
