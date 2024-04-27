import React, { useState } from "react";
import SuggestionContainer from "./SuggestionContainer";

const SearchBox = ({ val, setVal, label }) => {
  const [showSuggestions, setShowSuggestions] = useState(false);

  const handleChange = (event) => {
    setVal(event.target.value);
  };

  const handleFocus = () => {
    setShowSuggestions(true);
  };

  const handleBlur = () => {
    setTimeout(() => {
      setShowSuggestions(false);
    }, 200);
  };
  return (
    <>
      <div className="w-[20rem] h-[8.5rem] bg-custom-gray-1 border border-zinc-700 rounded-xl p-6 flex flex-col items-center">
        <p className="mb-4 text-center font-bold text-xl">{label}</p>
        <input
          type="text"
          value={val}
          onChange={handleChange}
          onFocus={handleFocus}
          onBlur={handleBlur}
          className="z-30 mb-2 w-full rounded-md border border-gray-400 p-1 text-black focus:outline-none"
        />
        {showSuggestions && <SuggestionContainer val={val} setVal={setVal} />}
      </div>
    </>
  );
};

export default SearchBox;
