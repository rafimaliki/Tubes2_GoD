import React, { useState } from "react";
import axios from "axios";

const SearchButton = ({ source, target, method, setResult }) => {
  const [loading, setLoading] = useState(false);

  const handleClick = () => {
    setResult({
      path: null,
      duration: null,
      error: null,
    });
    if (source && target && method) {
      setLoading(true);
      console.log(`Search ${method}`);
      console.log(`Source: ${source}`);
      console.log(`Target: ${target}`);

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
        });
    } else {
      console.log("Data belum lengkap");
    }
  };

  return (
    <>
      {loading && (
        <div className="fixed top-0 left-0 w-full h-full flex items-center justify-center bg-black bg-opacity-80 text-white z-50">
          <p className="text-2xl font-bold">FINDING PATH ...</p>
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
