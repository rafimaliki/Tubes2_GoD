import axios from "axios";
import React, { useState, useEffect } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";

import Solver from "./pages/Solver";
import About from "./pages/About";
import NotFound from "./pages/NotFound";
import NavBar from "./components/NavBar";

function App() {
  const [status, setStatus] = useState(0);

  useEffect(() => {
    axios
      .get("http://localhost:8080/api/status")
      .then((response) => {
        setStatus(response.data.status);
        console.log("Backend is running");
      })
      .catch((error) => {
        console.log("Backend is not running");
      });
  }, []);

  if (status === 0) {
    return <>Backend is not running!</>;
  } else {
    return (
      <Router>
        <div className="app-container min-h-screen">
          <NavBar />
          <Routes>
            <Route exact path="/" element={<Navigate to="/solver" />} />
            <Route path="/solver" element={<Solver />} />
            <Route path="/about" element={<About />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </div>
      </Router>
    );
  }
}

export default App;
