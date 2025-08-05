import React, { useState, useEffect } from "react";
import { toast } from "react-toastify";
import { getVideos } from "../services/videoService";
import VideoItem from "./VideoItem";
import "../css/videopage.css"

export default function Videos() {
    const [videos, setVideos] = useState([])
    const [page, setPage] = useState(1)
    const [loading, setLoading] = useState(false)

    useEffect(() => {
        setLoading(true);
        getVideos(page)
            .then(result => {
                if (result.data.videos.length === 0) {
                    setVideos([])
                } else {
                    setVideos(result.data.videos)
                }
            })
            .catch(err => {
                if (err.response) {
                    toast.error(err.response.data.message || "Failed to load videos. Please try again later.");
                } else if (err.request) {
                    toast.error("Network error. Please check your connection.");
                } else {
                    toast.error("An unexpected error occurred while loading videos.");
                }
            })
            .finally(() => {
                setLoading(false);
            });
    }, [page])

    const handlePreviousPage = () => {
        if (page > 1) {
            setPage(page - 1);
        }
    };

    const handleNextPage = () => {
        setPage(page + 1);
    };

    return (
        <div className="Video-container">
            <div className="videos-list">
                {loading ? (
                    <p style={{ color: "#bbb" }}>Loading videos...</p>
                ) : videos.length === 0 ? (
                    <p style={{ color: "#bbb" }}>No videos found.</p>
                ) : (
                    videos.map(video => (
                        <VideoItem key={video.Id} video={video} />
                    ))
                )}
            </div>

            <div className="pagination-container">
                <button
                    className="pagination-button"
                    onClick={handlePreviousPage}
                    disabled={page <= 1}
                >
                    Previous
                </button>
                <span className="page-indicator">Page {page}</span>
                <button
                    className="pagination-button"
                    onClick={handleNextPage}
                    disabled={videos.length === 0}
                >
                    Next
                </button>
            </div>
        </div>
    )
}