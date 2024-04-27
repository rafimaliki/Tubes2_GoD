import React from "react";

const Suggestion = ({ title, setVal }) => {
  const handleClick = () => {
    setVal(title);
    // console.log(title);
  };
  return (
    <div className="rounded-sm w-[17rem]">
      <button
        className="h-full rounded-lg w-full p-0 pl-1 text-gray-300 font-extralight hover:bg-zinc-800"
        onClick={handleClick}
      >
        {title}
      </button>
    </div>
  );
};

export default Suggestion;
