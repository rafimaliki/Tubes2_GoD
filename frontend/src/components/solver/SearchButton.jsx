import React from "react";
import axios from "axios";

const SearchButton = ({ source, target, method }) => {
  const HandleClick = () => {
    if (source && target && method) {
      console.log(`Search ${method}`);
      console.log(`Source: ${source}`);
      console.log(`Target: ${target}`);

      axios
        .get(
          `http://localhost:8080/api/${method}?source=${source}&target=${target}`
        )
        .then((response) => {
          console.log(
            response.data
            // `http://localhost:8080/api/${method}?source=${source}&target=${target}`
          );
          // console.log(response);
        })
        .catch((error) => {
          console.error(error);
        });
    } else {
      console.log("Data belum lengkap");
    }
  };

  return (
    <button
      className="bg-black w-[10.5rem] border border-gray-500 rounded-md py-1 font-bold
      hover:bg-gray-100 hover:text-black transition duration-200 ease-in-out"
      onClick={HandleClick}
    >
      GENERATE PATH
    </button>
  );
};

export default SearchButton;
