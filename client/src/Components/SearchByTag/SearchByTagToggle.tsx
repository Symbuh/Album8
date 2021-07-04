import React, {FC, useState} from 'react'
import apiInstance from '../../axiosConfig';

interface Props {
  /*
    We need our update function and that's it
  */
  setImagesByTag: () => void
  tags: any[]
}
/*
  The clicking the search by tag button should open a dropdown allowing the user the view the mapped tags and make a
  selection
*/
const SearchByTagsToggle: FC<Props> = ({ setImagesByTag, tags }) => {
  const [viewToggle, setViewToggle] = useState(false);

  const handleChange = (e: any) => {
    /*
      Save the selected tag and
      make a get request based on that tag,
      setting imagesByTag in the process
    */

    apiInstance.get(`/image/tag/${e.target.value}`)
    .then((res) => {
      console.log(`Recieved response from image/tag! ${res}`)
    })
    .catch((err) => console.log(err))
  }

  if (!viewToggle)  {
    return (
      <div>
        <button className="truthyButton" onClick={(e) => setViewToggle(false)}>All Photos</button>
        <button className="falsyButton" onClick={(e) => setViewToggle(true)}>Search By Tag</button>
      </div>
    )
  }
  return (
    <div>
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