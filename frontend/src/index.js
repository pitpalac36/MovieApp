import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import ErrorPage from './components/ErrorPage';
import Home from './components/Home';
import Movies from './components/Movies';
import Movie from './components/Movie';
import AddMovie from './components/AddMovie';
import ManageCatalogue from './components/ManageCatalogue';
import Login from './components/Login';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App/>,
    errorElement: <ErrorPage/>,
    children: [
      {index: true, element: <Home/>},
      {
        path: "/movies",
        element: <Movies/>
      },
      {
        path: "/movies/:id",
        element: <Movie/>
      },
      {
        path: "/admin/movie/0",
        element: <AddMovie/>
      },
      {
        path: "/manage-catalogue",
        element: <ManageCatalogue/>
      },
      {
        path: "/login",
        element: <Login/>
      },
    ]
  }
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>
);
