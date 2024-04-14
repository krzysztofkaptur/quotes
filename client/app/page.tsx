type Author = {
  id: number
  name: string
}

export default async function Home() {
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_BASE_API}/api/v1/authors`
  )
  const authors: Author[] = await response.json()

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <ul>
        {authors?.map(author => (
          <li key={author.id}>{author.name}</li>
        ))}
      </ul>
    </main>
  )
}
