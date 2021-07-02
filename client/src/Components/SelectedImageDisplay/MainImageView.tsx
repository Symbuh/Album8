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

const MainImageView: FC<Props> = ({id, name, url, description, tags, updateFunction}) => {

  const deleteImage = () => {
    apiInstance.delete(`/api/deleteimage/${id}`)
    .then((res) => {
      console.log(res)
    })
    .catch((err) => {
      console.log(err)
    })
    updateFunction()
  }

  if (id !== '' && name !== '' && url !== '' && tags) {
    return (
      <div id='selectedImageContainer'>
        <h2>{name}</h2>
        <img src={url} alt={name}/>
        <h4>{description}</h4>
        <div>
          Tags:
          {
            tags.map((tag: any) => {
              return (
                <div>
                  {tag}
                </div>
              )
            })
          }
        </div>
        <button onClick={deleteImage} />
      </div>
    )
  }
  return (<div id='selectedImageContainer'></div>)
}

export default MainImageView