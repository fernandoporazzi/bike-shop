import { MouseEvent, useState } from "react";

const Add = (): JSX.Element => {
  const [step, setStep] = useState(0);
  const [name, setName] = useState("");
  const [color, setColor] = useState("");
  const [stock, setStock] = useState(0);
  const [newBikeId, setNewBikeId] = useState<string | null>(null);

  const handleFormSubmit = async (event: MouseEvent<HTMLFormElement>) => {
    event.preventDefault();

    const body = {
      name,
      color,
      stock,
    };

    const res = await fetch("http://localhost:8080/bikes", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body),
    });
    const json = await res.json();

    setNewBikeId(json.id);
    setStep(1);
  };

  const handleImageChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
    event.preventDefault();

    if (!event.target.files) {
      return;
    }

    var formData = new FormData();
    formData.append("files", event.target.files[0]);

    const res = await fetch(`http://localhost:8080/bikes/upload/${newBikeId}`, {
      method: "POST",
      body: formData,
    });
    await res.json();
  };

  return (
    <div>
      {/* Basic bike informations */}
      {step === 0 && (
        <>
          <h2>Basic information about the bike</h2>
          <form method="POST" onSubmit={handleFormSubmit}>
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
        </>
      )}

      {/* Image upload */}
      {step === 1 && (
        <>
          <h2>Adding an image is optional</h2>
          <form>
            <input
              type="file"
              name="files"
              id="files"
              onChange={handleImageChange}
            />
          </form>
        </>
      )}
    </div>
  )
};

export default Add;
