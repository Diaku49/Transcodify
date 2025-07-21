import React, { useState, useEffect } from "react";
import { getVideos } from "../services/videoService";
import VideoItem from "./VideoItem";
import "../css/videopage.css"

export default function Videos() {
    const [videos, setVideos] = useState([])

    useEffect(() => {
        getVideos()
            .then(result => {
                setVideos(result.data)
            })
            .catch(err => {
                console.error(err)
            })
    }, [])

    return (
        <div className="Video-container">
            {videos.length === 0 ? (
                <p style={{ color: "#bbb" }}>No videos found.</p>
            ) : (
                videos.map(video => (
                    <VideoItem key={video.Id} video={video} />
                ))
            )
            }
        </div>
    )
}