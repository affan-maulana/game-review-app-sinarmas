export const routes = {
  games: {
    list: "/api/games",
    detail: (id: string) => `/api/games/${id}`,
    reviews: (id: string) => `/api/games/${id}/reviews`,
  },
};
