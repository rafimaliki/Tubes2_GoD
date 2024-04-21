import React, { useState } from "react";
import MethodToggle from "./MethodToggle";
import SearchButton from "./SearchButton";

const SearchBox = ({ source, target }) => {
  const [method, setMethod] = useState("");

  return (
    <div className="flex flex-col bg-custom-gray-1 h-full sm:w-[15rem] w-[20rem] rounded-xl border border-zinc-700">
      <p className="mt-6 text-center font-bold text-xl">METHOD</p>
      <div className="flex flex-col h-full justify-between">
        <div className="flex w-full justify-center mt-5">
          <MethodToggle label="BFS" method={method} setMethod={setMethod} />
          <MethodToggle label="IDS" method={method} setMethod={setMethod} />
        </div>
        <div className="flex items-center justify-center text-center mb-[2rem]">
          <SearchButton source={source} target={target} method={method} />
        </div>
      </div>
    </div>
  );
};
export default SearchBox;
