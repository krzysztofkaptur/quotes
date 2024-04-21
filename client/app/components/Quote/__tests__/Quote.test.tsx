import { expect, test, describe } from 'vitest'
import { render, screen } from '@testing-library/react'
import QuoteComp from '../'
import { Quote } from '../interface'

describe('Quote', () => {
  test('should render text and name from props', () => {
    const sampleText = "Smokes let's go"
    const sampleName = 'Ricky'

    const sampleQuote: Quote = {
      id: '123',
      text: sampleText,
      author_id: 1,
      name: sampleName
    }

    render(<QuoteComp quote={sampleQuote} />)
    expect(screen.getByText(sampleText)).toBeDefined()
    expect(screen.getByText(sampleName)).toBeDefined()
  })
})
