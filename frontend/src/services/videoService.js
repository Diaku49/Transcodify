import { api } from "./api"

const getVideos = (page) => api.get(`/video?page=${page}&limit=10`);
const getVideoById = (id) => api.get(`/video/${id}`);
const getVideoInfoById = (id) => api.get(`/video/${id}`);

const uploadVideo = (file, videoName, resolutions, token) => {
    const formData = new FormData();
    formData.append("file", file)

    // Create metadata object and send as JSON string
    const metadata = {
        VideoName: videoName,
        Resolutions: resolutions
    };
    formData.append("metadata", JSON.stringify(metadata));

    return api.post(`/video/upload`, formData, {
        headers: {
            "Content-Type": "multipart/form-data",
            "Authorization": `Bearer ${token}`,
        },
    });
}

export {
    getVideoById,
    getVideoInfoById,
    getVideos,
    uploadVideo
}