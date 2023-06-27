import { useEffect, useState } from "react";
import { useNavigate, useOutletContext, useParams } from "react-router-dom";
import Input from "./form/Input";
import Select from "./form/Select";
import TextArea from "./form/TextArea";

const AddMovie = () => {
    const navigate = useNavigate();
    const {jwtToken} = useOutletContext();

    const [error, setError] = useState();
    const [errors, setErrors] = useState([]);
    const { setAlertClassName } = useOutletContext();
    const { setAlertMessage } = useOutletContext()

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

    const [movie, setMovie] = useState({
        id: 0,
        title: "",
        release_date: "",
        runtime: "",
        mpaa_rating: "",
        description: "",
        genre: "",
    });

    let {id} = useParams();
    useEffect(() => {
        if (jwtToken == "") {
            navigate("/login");
            return;
        }
    }, [jwtToken, navigate]);

    const handleSubmit = (event) => {
        event.preventDefault();

        const requestOptions = {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(movie),
        }

        fetch(`/add`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    setAlertClassName("alert-danger");
                    setAlertMessage(data.message);
                } else {
                    setAlertClassName("d-none");
                    setAlertMessage("");
                    navigate("/movies");
                }
            })
            .catch(error => {
                setAlertClassName("alert-danger");
                setAlertMessage(error);
            })
    };

    const handleChange = () => (event) => {
        let value = event.target.value;
        let name = event.target.name;
        setMovie({
            ...movie,
            [name]: value
        });
    }

    return (
        <div>
            <h2>Add Movie</h2>
            <hr/>
            <form onSubmit={handleSubmit}>
                <input type="hidden" name="id" value={movie.id} id="id"></input>
                <Input
                    title={"Title"}
                    className={"form-control"}
                    type={"text"}
                    name={"title"}
                    value={movie.title}
                    onChange={handleChange("title")}
                    errorDiv={hasError("title")? "text-danger" : "d-none"}
                    errorMsg={"Please enter the title"}
                />

                <Input
                    title={"Release Date"}
                    className={"form-control"}
                    type={"date"}
                    name={"release_date"}
                    value={movie.release_date}
                    onChange={handleChange("date")}
                    errorDiv={hasError("release_date")? "text-danger" : "d-none"}
                    errorMsg={"Please enter a release date"}
                />

                <Input
                    title={"Runtime"}
                    className={"form-control"}
                    type={"text"}
                    name={"runtime"}
                    value={movie.runtime}
                    onChange={handleChange("runtime")}
                    errorDiv={hasError("runtime")? "text-danger" : "d-none"}
                    errorMsg={"Please enter a runtime"}
                />

                <Select
                    title={"MPAA Rating"}
                    name={"mpaa_rating"}
                    options={mpaa_options}
                    onChange={handleChange("mpaa_rating")}
                    placeholder={"Choose..."}
                    errorMsg={"Please choose"}
                    errorDiv={hasError("mpaa_rating") ? "text-danger" : "d-none"}
                />

                <TextArea
                    title="Description"
                    name={"description"}
                    value={movie.description}
                    rows={"3"}
                    onChange={handleChange("description")}
                    errorMsg={"Please enter a description"}
                    errorDiv={hasError("description") ? "text-danger" : "d-none"}
                />

                <Select
                    title={"Genre"}
                    name={"genre"}
                    options={genres_options}
                    value={movie.genre}
                    onChange={handleChange("genre")}
                    placeholder={"Choose..."}
                    errorMsg={"Please choose"}
                    errorDiv={hasError("genre") ? "text-danger" : "d-none"}
                />

                <input 
                    type="submit"
                    className="btn btn-primary"
                    value="Add"
                />
            </form>
        </div>
    )
}

export default AddMovie;