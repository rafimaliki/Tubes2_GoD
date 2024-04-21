import React from "react";

const SearchBox = ({ val, setVal, label }) => {
  const handleChange = (event) => {
    setVal(event.target.value);
    // console.log(event.target.value);
  };

  return (
    <>
      <div className="w-[20rem] bg-custom-gray-1 border border-zinc-700 rounded-xl p-6 flex flex-col justify-center items-center">
        <p className="mb-4 text-center font-bold text-xl">{label}</p>
        <input
          type="text"
          value={val}
          onChange={handleChange}
          className="mb-2 w-full rounded-md border border-gray-400 p-1 text-black focus:outline-none"
        />
      </div>
    </>
  );
};

export default SearchBox;
