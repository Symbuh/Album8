import React, {FC, useState, useEffect} from 'react'
import apiInstance from '../../axiosConfig';
import axios from 'axios'
interface Props {
  /*
    We need our update function and that's it
  */
  viewAllImages: () => void
  setImages: (data: any[]) => void
  tags: any[]
}
/*
  The clicking the search by tag button should open a dropdown allowing the user the view the mapped tags and make a
  selection
*/
const SearchByTagsToggle: FC<Props> = ({ viewAllImages ,setImages, tags }) => {
  const [viewToggle, setViewToggle] = useState(false);

  useEffect(() => {
    if (!viewToggle) {
      viewAllImages()
    }
  }, [viewToggle])

  const handleChange = (e: any) => {
    axios({
      method: 'get',
      url: `http://localhost:8080/api/image/tag/${e.target.value}`,
      headers: { },
      data : ''
    })
    .then(function (response: any) {
      const image = response.data
      setImages(image);
    })
    .catch(function (error: any) {
      console.log(error);
    });
  }

  if (!viewToggle || !tags)  {
    return (
      <div className="searchButtonsDiv">
        <button className="truthyButton" onClick={(e) => setViewToggle(false)}>All Photos</button>
        <button className="falsyButton" onClick={(e) => setViewToggle(true)}>Search By Tag</button>
      </div>
    )
  }
  return (
    <div className="searchButtonsDiv">
      <button className="falsyButton" onClick={(e) => setViewToggle(false)}>All Photos</button>
      <select onChange={handleChange}>
            {tags.map((tag) => {
              return (
                <option
                  key={tag}
                  value={tag}
                >
                  {tag}
                </option>
              );
            })}
          </select>
    </div>
  )
}

export default SearchByTagsToggle