import React, { useState, useEffect } from "react";
import axios from "axios";

function App() {
  const [message, setMessage] = useState("");
  const [power, setPower] = useState("");
  const [result, setResult] = useState("");

  // check api response
  useEffect(() => {
    axios
      .get("http://localhost:8080/api/data")
      .then((response) => {
        setMessage(response.data.message);
      })
      .catch((error) => {
        console.error("Error fetching datas:", error);
        setMessage("Error fetching data");
      });
  }, []);

  // testing function calc power of two
  const calculatePowerOfTwo = () => {
    axios
      .get(`http://localhost:8080/api/power-of-two?power=${power}`) // contoh pengiriman parameter fungsi ke backend
      // contoh multiple parameter 'http://localhost:8080/api/power-of-two?base=2&exponent=${power}'
      .then((response) => {
        setResult(response.data.result);
      })
      .catch((error) => {
        console.error("Error calculating power:", error);
        setResult("Error calculating power");
      });
  };

  return (
    <div>
      <h1>{message}</h1>
      <h1 className="text-red-500">Calculate Power of Two</h1>
      <div>
        <label htmlFor="powerInput">Enter power:</label>
        <input
          type="string"
          id="powerInput"
          value={power}
          onChange={(e) => setPower(e.target.value)}
        />
        <button onClick={calculatePowerOfTwo}>Calculate</button>
      </div>
      {result !== "" && <h2>Result: {result}</h2>}
    </div>
  );
}

export default App;
