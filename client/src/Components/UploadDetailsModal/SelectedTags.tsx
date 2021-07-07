import React, { FC } from 'react'

interface tagProps {
  tags: string[]
}

const SelectedTags: FC<tagProps> = ({ tags }) => {

  return (
    <div>
      <div id='tagsLabelUploadModal'>
        Tags:
      </div>
      {tags.map((tag) => {
        return (
          <div key={tag}>
            {tag}
          </div>
        )
      })}
    </div>
  )
}

export default SelectedTags