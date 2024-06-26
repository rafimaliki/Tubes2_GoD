import React, { useState } from "react";
import axios from "axios";

const SearchButton = ({ source, target, method, setResult }) => {
  const [loading, setLoading] = useState(false);
  const [elapsedTime, setElapsedTime] = useState(0);

  const handleClick = () => {
    setResult({
      path: [],
      duration: 0,
      searched: 0,
    });
    if (source && target && method) {
      setLoading(true);
      setElapsedTime(0);

      const intervalId = setInterval(() => {
        setElapsedTime((prevElapsedTime) => prevElapsedTime + 0.1);
      }, 100);

      axios
        .get(
          `http://localhost:8080/api/${method}?source=${source}&target=${target}`
        )
        .then((response) => {
          setResult(response.data);
          console.log(response.data);
        })
        .catch((error) => {
          console.error(error);
        })
        .finally(() => {
          setLoading(false);
          clearInterval(intervalId);
        });
    }
  };

  return (
    <>
      {loading && (
        <div className=" flex flex-col text-2xl font-bold fixed top-0 left-0 w-full h-full items-center justify-center bg-black bg-opacity-80 text-white z-50">
          <p>Finding Path ... </p>
          <p>{elapsedTime.toFixed(1)}s</p>
        </div>
      )}
      <button
        className="bg-black w-[10.5rem] border border-gray-500 rounded-md py-1 font-bold
        hover:bg-gray-100 hover:text-black transition duration-200 ease-in-out"
        onClick={handleClick}
      >
        GENERATE PATH
      </button>
    </>
  );
};

export default SearchButton;
