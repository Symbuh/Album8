import React, {FC, useState} from 'react'

interface Props {
  /*
    We need our update function and that's it
  */
  setImagesByTag: () => void
  tags: any[]
}

const SearchByTagsToggle: FC<Props> = ({ setImagesByTag, tags }) => {
  return (
    <div>

    </div>
  )
}

export default SearchByTagsToggle