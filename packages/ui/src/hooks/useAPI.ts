import { useEffect, useState } from "react";
import { BASE_API_URL } from "../common/constants";

interface UseResource<T> {
  data: T;
  loading: boolean;
  error: Error | null;
}

export default function useAPI<T>(route: string): UseResource<T> {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const [data, setData] = useState<T | null>(null);

  useEffect(() => {
    async function fetchData() {
      try {
        const res = await fetch(`${BASE_API_URL}${route}`);
        const jsonData = await res.json();
        setData(jsonData);
      } catch (e) {
        setError(e);
      }
      setLoading(false);
    }
    fetchData();
  }, [route]);

  return {
    // If loading is false and error is false data can't be null so we're
    // asserting non-null for TypeScript compilation.
    data: data!,
    loading,
    error,
  };
}
