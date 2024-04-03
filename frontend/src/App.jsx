import React, { useState, useEffect } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import axios from "axios";

import Landing from "./pages/Landing";
import Game from "./pages/Game";
import About from "./pages/About";
import NotFound from "./pages/NotFound";
import NavBar from "./components/NavBar";

function App() {
  const [status, setStatus] = useState(0);

  /* Check if the backend is running */
  useEffect(() => {
    axios
      .get("http://localhost:8080/api/data")
      .then((response) => {
        setStatus(response.data.status);
        console.log("Backend is running");
      })
      .catch((error) => {
        console.error("Error fetching datas:", error);
      });
  }, []);

  if (status === 0) {
    return <>Backend is not running!</>;
  } else {
    return (
      <>
        <Router>
          <NavBar />
          <Routes>
            <Route exact path="/" element={<Landing />} />
            <Route path="/game" element={<Game />} />
            <Route path="/about" element={<About />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </Router>
      </>
    );
  }
}

export default App;
