import React from "react";

function CreateSpeaker() {
  return (
    <div>
      <h3>Create a speaker</h3>
      <form>
        <section>
          <label htmlFor="name">Name</label>
          <input type="text" name={"name"} />
        </section>
        <section>
          <label htmlFor="bio">Bio</label>
          <input type="text" name={"bio"} />
        </section>
        <section>
          <label htmlFor="speakerName">Headline</label>
          <input type="text" name={"headline"} />
        </section>
        <section>
          {/*For now it's just a url, later we might add a photo upload*/}
          <label htmlFor="speakerName">Photo</label>
          <input type="text" name={"speakerName"} />
        </section>
        <button type={"submit"}>Create</button>
      </form>
    </div>
  );
}

export default CreateSpeaker;
