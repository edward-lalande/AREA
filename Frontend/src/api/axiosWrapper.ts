import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";

class AxiosWrapper {

  	private axiosInstance: AxiosInstance;

    constructor(baseURL: string) {

        this.axiosInstance = axios.create({
            baseURL,
            headers: {
                "Content-Type": "application/json"
            }
        });

        this.axiosInstance.interceptors.response.use(
            (response: AxiosResponse) => response,
            (error) => {
                console.error("API Error:", error.response || error.message);
                return Promise.reject(error);
            }
        );

  }

    public async get<T>(
        serviceRoute: string,
        endpoint: string,
        params?: Record<string, any>,
        config?: AxiosRequestConfig
    ): Promise<T> {
        return this.request<T>("GET", serviceRoute, endpoint, undefined, {
            ...config,
            params
        });
    }

    public async post<T>(
        serviceRoute: string,
        endpoint: string,
        data?: Record<string, any>,
        config?: AxiosRequestConfig
    ): Promise<T> {
        return this.request<T>("POST", serviceRoute, endpoint, data, config);
    }

    public async put<T>(
        serviceRoute: string,
        endpoint: string,
        data?: Record<string, any>,
        config?: AxiosRequestConfig
    ): Promise<T> {
        return this.request<T>("PUT", serviceRoute, endpoint, data, config);
    }

    public async delete<T>(
        serviceRoute: string,
        endpoint: string,
        config?: AxiosRequestConfig
    ): Promise<T> {
        return this.request<T>("DELETE", serviceRoute, endpoint, undefined, config);
    }

    private async request<T>(
        method: "GET" | "POST" | "PUT" | "DELETE",
        serviceRoute: string,
        endpoint: string,
        data?: Record<string, any>,
        config?: AxiosRequestConfig
    ): Promise<T> {
        const url = `${serviceRoute}${endpoint}`;
        const requestConfig: AxiosRequestConfig = {
            ...config,
            method,
            url,
            data
        };

        const response = await this.axiosInstance.request<T>(requestConfig);
        return response.data;
    }

}

export default AxiosWrapper;
