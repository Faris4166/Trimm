"use client";
import { useState } from "react";
import axios from "axios";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function Home() {
  const [url, setUrL] = useState("");
  const [shortUrl, setShortUrl] = useState("");

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    try {
      const response = await axios.post("/shorten", { url });
      const { short_url } = response.data;
      setShortUrl(short_url);
    } catch (error) {
      console.error(error);
    }
  };
  const handleOpenUrl = () => {
    if (shortUrl) {
      window.open(`http://localhost:8080/${shortUrl}`, "_blank");
    }
  };
  const handCopyUrl = () => {
    if (shortUrl) {
      navigator.clipboard.writeText(`http://localhost:8080/${shortUrl}`);
    }
  };
  return (
    <div className="felx flex-col gap-2.5 justify-center items-center p-2.5">
      <h1 className="text-4xl">Shorten URL</h1>
      <form onSubmit={handleSubmit} className="flex flex-col gap-2.5">
        <Input
          placeholder="Enter URL"
          type="text"
          value={url}
          onChange={(event) => setUrL(event.target.value)}
        />
        <Button type="submit">Shortend</Button>
      </form>
      {shortUrl && (
        <div>
          <div>
            <p>
              Shortend URL:{" "}
              <a
                href={`http://localhost:8080/${shortUrl}`}
                target="_black"
                rel="noopener norferrer"
              >{`http://localhost:8080/${shortUrl}`}</a>
            </p>
          </div>
          <Button onClick={handleOpenUrl}>Open URL</Button>
          <Button onClick={handCopyUrl}>Copy URL</Button>
        </div>
      )}
    </div>
  );
}
