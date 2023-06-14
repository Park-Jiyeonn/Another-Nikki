from fastapi import FastAPI, Request, HTTPException
from fastapi.middleware.cors import CORSMiddleware
import os

app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.post("/github_webhook")
async def process_data1(req: Request):
    headers = req.headers
    event_type = headers.get('X-GitHub-Event', None)
    if event_type == 'push':
        dir = os.getcwd()

        os.system("git pull --recurse-submodules")
        os.system("lsof -i tcp:8888 | awk 'NR!=1 {print $2}' | xargs kill")

        os.chdir(dir + "/Another-Nikki-Server")     # 重启后端
        os.system("go run . &")

        os.chdir(dir + "/Another-Nikki-Web")        # 编译前端
        os.system("yarn run build")

        os.chdir(dir)       # 切换回去

        return {
            "msg":"success",
        }
    else:
        raise HTTPException(status_code=400, detail="not push event")

if __name__ == "__main__":
    port = 8849
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=port)