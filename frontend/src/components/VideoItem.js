import React from "react";
import '../css/videoItem.css';

export default function VideoItem({ video }) {
    return (
        <div className="video-item">
            <h3 className="video-titel">{video.name}</h3>
            {video.VideoVariant && video.VideoVariant.length > 0 ? (
                <div className="variant-list">
                    {video.VideoVariant.map(variant => (
                        <div className="variant-item" key={variant.id}>
                            <span className="variant-resolution">{variant.resolution}:</span>
                            <a
                                href={variant.url}
                                target="_blank"
                                rel="noopener noreferrer"
                                className="variant-link"
                            >
                                Download
                            </a>
                        </div>
                    ))}
                </div>
            ) : (
                <div className="no-variants">No video variants available.</div>
            )}
        </div>
    );
}