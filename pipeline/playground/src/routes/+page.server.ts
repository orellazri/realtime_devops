export const actions = {
  start: async ({ request }) => {
    const data = await request.formData();
    console.log(data);
  }
};
