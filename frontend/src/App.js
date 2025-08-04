import "./App.css"
import 'react-toastify/dist/ReactToastify.css'
import { ToastContainer } from 'react-toastify'
import Header from "./components/common/Header";
import Footer from "./components/common/Footer";
import { Route, Routes } from "react-router-dom";
import HomePage from "./components/pages/HomePage";
import UploadPage from "./components/pages/UploadPage";
import Video from "./components/Videos";
import AuthPage from "./components/pages/AuthPage";
import { useState } from "react";

function App() {

  const [isLoggedIn, setIsLoggedIn] = useState(false)

  return (
    <div className="App">
      <div className="container">
        <Header isLoggedIn={isLoggedIn} setIsLoggedIn={setIsLoggedIn}></Header>

        <ToastContainer position="top-center" theme="dark" />

        <main className="main-content">
          <Routes>
            <Route path="/" element={<HomePage></HomePage>} />
            <Route path="/video/upload" element={<UploadPage></UploadPage>} />
            <Route path="/video" element={<Video></Video>} />
            <Route path="/auth" element={<AuthPage setIsLoggedIn={setIsLoggedIn} ></AuthPage>} />
            <Route path="*" element={<></>} />
          </Routes>
        </main>

        <Footer></Footer>
      </div>
    </div>
  );
}

export default App;
