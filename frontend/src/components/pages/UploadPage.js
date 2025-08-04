import React, { useEffect, useState } from "react";
import styles from "../../css/uploadpage.module.css"
import { getVideoInfoById, uploadVideo } from "../../services/videoService";

export default function UploadPage() {
    const [videoFile, setVideoFile] = useState(null)
    const [videoName, setVideoName] = useState("");
    const [resolutions, setResolutions] = useState(["1080"])
    const [isUploading, setIsUploading] = useState(false)
    const [uploadStatus, setUploadStatus] = useState(null)
    const [uploadId, setUploadId] = useState(null)
    const [isTranscoding, setIsTranscoding] = useState(false)
    const [transcodeStatus, setTranscodeStatus] = useState(null)

    useEffect(() => {
        const trStatus = localStorage.getItem("transcodeStatus")
        if (trStatus) {
            setIsTranscoding(true)
            setIsUploading(false)
            setTranscodeStatus(trStatus)
            const videoId = localStorage.getItem("videoId")
            setUploadId(videoId)
            setUploadStatus("Uploaded successfully")
        }
    }, [uploadId, isTranscoding, isUploading, transcodeStatus])

    const handleSubmit = async (e) => {
        e.preventDefault()

        if (!videoFile) {
            alert("Please select a video file")
            return;
        }

        setIsUploading(true);
        setUploadStatus("Uploading video...")

        try {
            const response = await uploadVideo(videoFile, videoName, resolutions)
            if (response.status === 200) {
                setUploadStatus(response.data.message)
                setUploadId(response.data.id)
                localStorage.setItem("videoId", response.data.id)
                localStorage.setItem("transcodeStatus")
                setVideoFile(null)
                setVideoName(null)
                setResolutions([])
            } else {
                setUploadStatus(response.data.message)
            }
        } catch (err) {
            console.log(err)
            setUploadStatus("An error occured while uploading")
        }

        setIsUploading(false)
    }

    const handleVideoNameChange = (e) => {
        setVideoName(e.target.value);
    };

    const handleResolutionChange = (e) => {
        const value = e.target.value;
        setResolutions((prev) =>
            prev.includes(value)
                ? prev.filter((resolution) => resolution !== value)
                : [...prev, value]
        )
    }
    const handleFileChange = (e) => {
        setVideoFile(e.target.files[0])
    }
    const handleGetInfo = async (e) => {
        e.preventDefault()

        if (!uploadId) {
            alert("no uploaded video for transcoding");
            return;
        }
        setIsUploading(false)
        setIsTranscoding(true)
        setTranscodeStatus("Slave starting...")
        let response

        try {
            response = await getVideoInfoById(uploadId)
            if (response.status === 200) {
                setTranscodeStatus(response.data.message)
            } else {
                setTranscodeStatus(response.data.message)
            }
        } catch (err) {
            setTranscodeStatus("Couldn get info about the transcode progress")
        }

        if (response.data.message === "Finished") {
            setIsTranscoding(false)
            setUploadId(null)
        }
    }

    return (
        <div className={styles.uploadContainer}>
            <h1>Upload Your Video</h1>

            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="videoName">Video Name:</label>
                    <input
                        type="text"
                        id="videoName"
                        value={videoName}
                        onChange={handleVideoNameChange}
                        placeholder="Enter a name for your video"
                        required
                    />
                </div>

                <input
                    type="file"
                    accept="video/*"
                    onChange={handleFileChange}
                    required
                />

                <div>
                    <label>Select Resolutions:</label>
                    <label>
                        <input
                            type="checkbox"
                            value="1080p"
                            checked={resolutions.includes("1080p")}
                            onChange={handleResolutionChange}
                        />
                        1080p
                    </label>
                    <label>
                        <input
                            type="checkbox"
                            value="720p"
                            checked={resolutions.includes("720p")}
                            onChange={handleResolutionChange}
                        />
                        720p
                    </label>
                    <label>
                        <input
                            type="checkbox"
                            value="480p"
                            checked={resolutions.includes("480p")}
                            onChange={handleResolutionChange}
                        />
                        480p
                    </label>
                    <label>
                        <input
                            type="checkbox"
                            value="360p"
                            checked={resolutions.includes("360p")}
                            onChange={handleResolutionChange}
                        />
                        360p
                    </label>
                </div>

                <button type="submit" disabled={isUploading}>
                    {isUploading ? "Uploading..." : "Upload Video"}
                </button>
            </form>

            <button onClick={handleGetInfo} disabled={!uploadId || isTranscoding}>
                Get Upload Info
            </button>

            {uploadStatus && <p>{uploadStatus}</p>}
            {transcodeStatus && <p>{transcodeStatus}</p>}
        </div>
    )
}