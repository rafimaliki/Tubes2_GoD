import React from "react";
import axios from "axios";

const SearchButton = ({ source, target }) => {
  const HandleClick = () => {
    console.log(`Source: ${source}`);
    console.log(`Target: ${target}`);

    axios
      .get(`http://localhost:8080/api/search?source=${source}&target=${target}`)
      .then((response) => {
        console.log(response);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <button
      className="m-5 border border-zinc-700 px-12 py-1 rounded-sm transition duration-300 ease-in-out hover:bg-zinc-900 hover:border-zinc-900 "
      onClick={HandleClick}
    >
      Search
    </button>
  );
};

export default SearchButton;
