import QuoteComp from './components/Quote'
import { fetchRandomQuote } from './services/quotes'

import type { Quote } from './components/Quote/interface'

export const revalidate = 0

export default async function Home() {
  const response = await fetchRandomQuote()
  const quote: Quote = await response.json()

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <QuoteComp quote={quote} />
    </main>
  )
}
