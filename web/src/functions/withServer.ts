export const withServer = (path: string): string => `${import.meta.env.VITE_API_BASE_URL}${path}`
