import { useEffect, useState, MouseEvent } from "react";
import { useParams } from "react-router-dom";

type Params = {
  id: string;
}

type Response = {
  name: string;
  color: string;
  stock: number;
}

const Bike = (): JSX.Element => {
  let { id } = useParams<Params>();

  const [name, setName] = useState("");
  const [color, setColor] = useState("");
  const [stock, setStock] = useState(0);

  useEffect(() => {
    const request = async () => {
      const res = await fetch(`http://localhost:8080/bikes/${id}`);
      const data: Response = await res.json();

      setName(data.name);
      setColor(data.color);
      setStock(data.stock);
    };

    request();

  }, [id]);

  const handleFormSubmit = async (event: MouseEvent<HTMLFormElement>) => {
    event.preventDefault();

    const body = {
      id,
      name,
      color,
      stock,
    };

    const res = await fetch(`http://localhost:8080/bikes/${id}`, {
      method: "PUT",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body),
    });

    await res.json();

    alert("Updated");
  };

  return (
    <div>
      <h2>Update bike details</h2>
      <form onSubmit={handleFormSubmit}>
        <div>
          <label htmlFor="name">Bike name</label>
          <input
            type="text"
            name="name"
            id="name"
            required
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>

        <div>
          <label htmlFor="color">Bike color</label>
          <input
            type="text"
            name="color"
            id="color"
            required
            value={color}
            onChange={(e) => setColor(e.target.value)}
          />
        </div>

        <div>
          <label htmlFor="stock">Available stock</label>
          <input
            type="number"
            name="stock"
            id="stock"
            required
            value={stock}
            onChange={(e) => setStock(+e.target.value)}
          />
        </div>

        <button type="submit">Save</button>
      </form>
    </div>
  )
};

export default Bike;
