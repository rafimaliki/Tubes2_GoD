import React from "react";

const Result = ({ id, title, url }) => {
  return (
    <a
      href={url}
      target="_blank"
      rel="noopener noreferrer"
      className="p-0 w-full mt-4 flex items-center bg-custom-gray-1 h-fit rounded-xl border border-zinc-700 hover:border-white transition duration-200 ease-in-out hover:ml-6"
    >
      <p className="flex justify-center items-center text-center h-[3rem] text-2xl sm:text-3xl lg:text-4xl font-bold px-4 border-zinc-700 hover:border-white whitespace-normal truncate lg:whitespace-nowrap">
        {id}
      </p>
      <p className="font-light text-lg sm:text-xl lg:text-2xl whitespace-normal truncate">
        {title}
      </p>
    </a>
  );
};

export default Result;
