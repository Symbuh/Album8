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
      <div id='selectedImageContainer'>
        <h2>{name}</h2>
        <img src={url} alt={name}/>
        <h4>{description}</h4>
        <div>
          Tags:
          {
            tags !== null && tags.map((tag: any) => {
              return (
                <div>
                  {tag}
                </div>
              )
            })
          }
        </div>
        <div>
          <button onClick={handleClick}>Delete Image</button>
        </div>
      </div>
    )
  }
  return (<div id='selectedImageContainer'></div>)
}

export default MainImageView