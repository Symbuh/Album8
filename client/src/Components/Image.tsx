import React, { FC } from 'react';

interface ImageProps {
    id: string
    name: string
    url: string
    description: string
    tags?: string[]
}


const Image: FC<ImageProps> = ({id, url, name, description, tags}) => {
  return (
    <div key={id} className='singleImageContainer'>
      <div className='imageName'>
        Name: {name}
      </div>
      <div className='imageContainer'>
        <img src={url} alt={name}></img>
      </div>
      <div className='description'>
        Description: {description}
      </div>
    </div>
  )
}

export default Image;