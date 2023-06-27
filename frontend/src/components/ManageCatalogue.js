import { useEffect, useState } from "react";
import { useNavigate, useOutletContext } from "react-router-dom";
import Input from "./form/Input";
import Select from "./form/Select";
import TextArea from "./form/TextArea";

const ManageCatalogue = () => {
    const navigate = useNavigate();
    const {jwtToken} = useOutletContext();
    const [movies, setMovies] = useState([]);

    const [error, setError] = useState();
    const [errors, setErrors] = useState([]);
    const { setAlertClassName } = useOutletContext();
    const { setAlertMessage } = useOutletContext();

    const mpaa_options = [
        {id: "G", value: "G"},
        {id: "PG", value: "PG"},
        {id: "PG-13", value: "PG-13"},
        {id: "R", value: "R"},
        {id: "18A", value: "18A"},
    ];

    const genres_options = [
        {id: "14", value: "Comedy"},
        {id: "15", value: "Horror"},
        {id: "16", value: "Sci-Fi"},
        {id: "17", value: "Romance"},
        {id: "18", value: "Action"},
    ];

    const hasError = (key) => {
        return errors.indexOf(key) !== -1;
    }

    useEffect(() => {
        if (jwtToken == "") {
            navigate("/login");
            return;
        }
    }, [jwtToken, navigate]);

    useEffect(() => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        };

        fetch(`/movies`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
            setMovies(data);
        })
        .catch(err => {
            console.log(err);
        })
    }, []);

    const deleteMovie = (movie) => {
      const requestOptions = {
        method: "DELETE",
      };

      fetch(`/delete/movie/${movie.id}`, requestOptions)
        .then((id) => {
          setAlertClassName("d-none");
          setAlertMessage("");
          fetch(`/movies`, {
            method: "GET",
            headers: { "Content-Type": "application/json" },
          })
            .then((response) => response.json())
            .then((data) => {
              setMovies(data);
            })
            .catch((err) => {
              console.log(err);
            });
        })
        .catch((error) => {
          setAlertClassName("alert-danger");
          setAlertMessage(error);
        });
    };

    const handleChange = () => (event) => {
    }

    return (
        <div>
            <h2>Manage Catalogue</h2>
            <hr/>
            <div>
            {movies.map((m) => (
                        <div key={m.id}>
                        <input type="hidden" name="id" value={m.id} id="id"></input>
                        <Input
                            title={"Title"}
                            className={"form-control"}
                            type={"text"}
                            name={"title"}
                            value={m.title}
                            onChange={handleChange("title", m)}
                            errorDiv={hasError("title")? "text-danger" : "d-none"}
                            errorMsg={"Please enter the title"}
                        />
        
                        <Input
                            title={"Release Date"}
                            className={"form-control"}
                            type={"date"}
                            name={"release_date"}
                            value={m.release_date.split("T")[0]}
                            onChange={handleChange("date", m)}
                            errorDiv={hasError("release_date")? "text-danger" : "d-none"}
                            errorMsg={"Please enter a release date"}
                        />
        
                        <Input
                            title={"Runtime"}
                            className={"form-control"}
                            type={"text"}
                            name={"runtime"}
                            value={m.runtime}
                            onChange={handleChange("runtime", m)}
                            errorDiv={hasError("runtime")? "text-danger" : "d-none"}
                            errorMsg={"Please enter a runtime"}
                        />
        
                        <Select
                            title={"MPAA Rating"}
                            name={"mpaa_rating"}
                            value={m.mpaa_rating}
                            options={mpaa_options}
                            onChange={handleChange("mpaa_rating", m)}
                            placeholder={"Choose..."}
                            errorMsg={"Please choose"}
                            errorDiv={hasError("mpaa_rating") ? "text-danger" : "d-none"}
                        />
        
                        <TextArea
                            title="Description"
                            name={"description"}
                            value={m.description}
                            rows={"3"}
                            onChange={handleChange("description", m)}
                            errorMsg={"Please enter a description"}
                            errorDiv={hasError("description") ? "text-danger" : "d-none"}
                        />
        
                        <Select
                            title={"Genre"}
                            name={"genre"}
                            options={genres_options}
                            value={m.Genre.genre}
                            onChange={handleChange("genre", m)}
                            placeholder={m.Genre.genre}
                            errorMsg={"Please choose"}
                            errorDiv={hasError("genre") ? "text-danger" : "d-none"}
                        />

                        <button 
                            className="btn btn-primary"
                            value="Delete"
                            onClick={() => deleteMovie(m)}
                            >Delete</button>
                        <hr/>
                    </div>
                    
                    ))}
            </div>
            
        </div>
    )
}

export default ManageCatalogue;