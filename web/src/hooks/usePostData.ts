import { useState } from 'react';

interface PostOptions<T> {
    method?: 'POST' | 'PUT' | 'PATCH';
    headers?: HeadersInit;
    body: T;
}

export const usePostData = <T, R>(url: string, options: PostOptions<T>) => {
    const [data, setData] = useState<R | null>(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const postData = async () => {
        setLoading(true);
        try {
            const response = await fetch(url, {
                method: options.method || 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    ...options.headers,
                },
                body: JSON.stringify(options.body),
            });
            const responseData: R = await response.json();
            setData(responseData);
        } catch (error) {
            setError(`An error occurred: ${error}`);
        } finally {
            setLoading(false);
        }
    };

    return { data, loading, error, postData };
};