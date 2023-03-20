import { json } from "@sveltejs/kit";

export async function POST({ request }) {
  const { services } = await request.json();
  console.log(services);

  return json({ a: "b" });
}
