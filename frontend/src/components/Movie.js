import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const Movie = () => {

    const [movie, setMovie] = useState();
    let { id } = useParams();

    useEffect(() => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        };

        fetch(`/movie/${id}`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
            setMovie(data);
        })
        .catch(err => {
            console.log(err);
        })
    }, []);

    return (
        movie &&
        <div>
            <h2>Movie: {movie.title}</h2>
            <small><em>{movie.release_date}, {movie.runtime} minutes, Rated {movie.mpaa_rating}</em></small>
            <hr/>
            <p>{movie.description}</p>
            {(movie.image !== "") && <img style={{ width: 500, height: 600 }} src={`data:image/jpeg;base64,${movie.image}`}/>}
        </div>
    )
}

export default Movie;