import { api } from "./api"

const getVideos = () => api.get("/video");
const getVideoById = (id) => api.get(`/video/${id}`);
const getVideoInfoById = (id) => api.get(`/video/${id}`);

const uploadVideo = (file, videoName, resolutions) => {
    const formData = new FormData();
    formData.append("file", file)
    formData.append("resolutions", JSON.stringify(resolutions));
    formData.append("videoName", videoName);

    return api.post(`/video/upload`, formData, {
        headers: {
            "Content-Type": "multipart/form-data",
        },
    });
}

export {
    getVideoById,
    getVideoInfoById,
    getVideos,
    uploadVideo
}