import React from "react";

const InputBox = ({ val, setVal, label }) => {
  const handleChange = (event) => {
    setVal(event.target.value);
    // console.log(event.target.value);
  };

  return (
    <>
      <div className="w-[18rem] border border-zinc-700 rounded-sm p-4 flex flex-col justify-center items-center mx-3">
        <p className="mb-2 text-center font-bold">{label}</p>
        <input
          type="text"
          value={val}
          onChange={handleChange}
          className="w-full rounded-md border border-gray-400 p-1 text-black focus:outline-none"
        />
      </div>
    </>
  );
};

export default InputBox;
