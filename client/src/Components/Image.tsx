import React, { FC } from 'react';

interface ImageProps {
  id: number
  name: string
  url: string
  description: string
  tags?: string[]
}

const Image: FC<ImageProps> = ({id, name, url, description, tags}) => {
  return (
    <div key={id}>
      <div className='imageName'>
        {name}
      </div>
      <div className='imageContainer'>
        <img src={url} alt={name}></img>
      </div>
      <div className='description'>
        {description}
      </div>
    </div>
  )
}

export default Image;