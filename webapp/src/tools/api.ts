function getApiURL() {
  return localStorage.getItem('apiURL') || import.meta.env.VITE_API_URL || 'http://127.0.0.1:8080';
}

export async function apiGet(uri: string) {
  const URL = getApiURL();

  if (!URL) {
    console.log("Failed to send request because API_URL is not set");
    return
  }

  if (uri.startsWith('/')) {
    uri = uri.substring(1);
  }

  const headers = {
    'Accept': 'application/json'
  }

  const response = await fetch(`${URL}/${uri}`,
    {
      method: "get",
      headers: headers
    }
  );

  return response;
};

export async function apiPost(uri: string, payload: object) {
  const URL = getApiURL();

  if (!URL) {
    console.log("Failed to send request because API_URL is not set");
    return
  }

  if (uri.startsWith('/')) {
    uri = uri.substring(1);
  }

  const headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
  }

  const response = await fetch(`${URL}/${uri}`,
    {
      method: "post",
      headers: headers,
      body: JSON.stringify(payload)
    }
  );

  return response; 
};

export async function apiPut(uri: string, payload: object) {
  const URL = getApiURL();

  if (!URL) {
    console.log("Failed to send request because API_URL is not set");
    return
  }

  if (uri.startsWith('/')) {
    uri = uri.substring(1);
  }

  const headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
  }

  const response = await fetch(`${URL}/${uri}`,
    {
      method: "put",
      headers: headers,
      body: JSON.stringify(payload)
    }
  );

  return response; 
};

export async function apiDelete(uri: string) {
  const URL = getApiURL();

  if (!URL) {
    console.log("Failed to send request because API_URL is not set");
    return
  }

  if (uri.startsWith('/')) {
    uri = uri.substring(1);
  }

  const headers = {
    'Accept': 'application/json'
  }

  const response = await fetch(`${URL}/${uri}`,
    {
      method: "delete",
      headers: headers
    }
  );

  return response; 
}