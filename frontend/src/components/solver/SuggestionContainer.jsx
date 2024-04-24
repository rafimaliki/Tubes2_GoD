import React, { useEffect } from "react";
import Suggestion from "./Suggestion";
import axios from "axios";

const SuggestionContainer = ({ val, setVal }) => {
  // useeffect on variable val
  const [response, setResponse] = React.useState([]);

  useEffect(() => {
    if (val) {
      const queryParams = {
        action: "query",
        format: "json",
        gpssearch: val,
        generator: "prefixsearch",
        prop: "pageprops|pageimages|pageterms",
        redirects: "",
        ppprop: "displaytitle",
        piprop: "thumbnail",
        pithumbsize: "160",
        pilimit: "30",
        wbptterms: "description",
        gpsnamespace: 0,
        gpslimit: 5,
        origin: "*",
      };

      axios({
        method: "get",
        url: "https://en.wikipedia.org/w/api.php",
        params: queryParams,
        headers: {
          "Api-User-Agent": "13522137@std.stei.itb.ac.id",
        },
      })
        .then((response) => {
          setResponse(response.data.query.pages);
          console.log(response.data);
        })
        .catch((error) => {
          console.error(error);
        });
    }
  }, [val]);

  const renderSuggestion = () => {
    return Object.keys(response).map((pageId) => {
      const obj = response[pageId];
      return <Suggestion key={pageId} title={obj.title} setVal={setVal} />;
    });
  };

  return <div className=" flex flex-col z-10 -mt-2">{renderSuggestion()}</div>;
};

export default SuggestionContainer;
