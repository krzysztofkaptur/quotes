export const fetchRandomQuote = () => {
  return fetch(`${process.env.NEXT_PUBLIC_BASE_API}/api/v1/quotes/random`, {
    cache: 'no-store'
  })
}
