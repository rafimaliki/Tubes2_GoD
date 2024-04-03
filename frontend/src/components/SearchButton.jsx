import React from "react";
import axios from "axios";

const SearchButton = ({ source, target, endpoint, label }) => {
  const HandleClick = () => {
    console.log(`Source: ${source}`);
    console.log(`Target: ${target}`);

    axios
      .get(
        `http://localhost:8080/api/${endpoint}?source=${source}&target=${target}`
      )
      .then((response) => {
        console.log(response);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <button
      className="w-[18rem] mx-3 my-1 border border-zinc-700 px-12 py-1 rounded-sm transition duration-300 ease-in-out hover:bg-zinc-900 hover:border-zinc-900 "
      onClick={HandleClick}
    >
      {label}
    </button>
  );
};

export default SearchButton;
