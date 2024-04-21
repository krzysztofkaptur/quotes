import type { Props } from './interface'

export default function Quote({quote}: Props) {
  return (
    <div className="flex flex-col">
        <i>{quote.text}</i>
        <small className="self-end">- <span>{quote.name}</span></small>
      </div>
  )
}