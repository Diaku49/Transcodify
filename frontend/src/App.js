import "./App.css"
import Header from "./components/common/Header";
import Footer from "./components/common/Footer";
import { Route, Routes } from "react-router-dom";
import HomePage from "./components/pages/HomePage";
import UploadPage from "./components/pages/UploadPage";
import Video from "./components/Videos";
import AuthPage from "./components/pages/AuthPage";

function App() {

  return (
    <div className="App">
      <div className="container">
        <Header ></Header>


        <main className="main-content">
          <Routes>
            <Route path="/" element={<HomePage></HomePage>} />
            <Route path="/video/upload" element={<UploadPage></UploadPage>} />
            <Route path="/video" element={<Video></Video>} />
            <Route path="/auth" element={<AuthPage ></AuthPage>} />
            <Route path="*" element={<></>} />
          </Routes>
        </main>

        <Footer></Footer>
      </div>
    </div>
  );
}

export default App;
