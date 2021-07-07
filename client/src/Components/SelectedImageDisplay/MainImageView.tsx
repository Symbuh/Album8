import React, {FC, useState} from 'react'
import { apiInstance } from '../../axiosConfig'

interface Props {
  id: string
  name: string
  url: string
  description: string
  tags?: any
  updateFunction: () => void
}

// something
const MainImageView: FC<Props> = ({id, name, url, description, tags, updateFunction}) => {

  const deleteImage = () => {
    console.log('calling delete image')
    apiInstance.delete(`/api/deleteimage/${id}`)
    .then((res) => {
      console.log(res)
      updateFunction()
    })
    .catch((err) => {
      console.log(err)
    })
  }

  const handleClick = () => {
    deleteImage()
    updateFunction()
  }

  if (id !== '' && name !== '' && url !== '') {
    return (
      <div>

        <div id="mainImageHeaderContainer">
          <h2 id="mainImageName">{name}</h2>
          <button onClick={handleClick}>Delete Image</button>
        </div>
        <div id='selectedImageContainer'>
          <img id="mainImage" src={url} alt={name}/>
        </div>
        <h4>{description}</h4>
        <a href={url}>Permalink URL</a>
        <div className="tagsContainer">
          Tags:
          {
            tags !== null && tags !== undefined && tags.map((tag: any) => {
              return (
                <div>
                  {tag}
                </div>
              )
            })
          }
        </div>
        <div>

        </div>
      </div>
    )
  }
  return (<div id='selectedImageContainer'></div>)
}

export default MainImageView