package xavier.ricardo.paulobet.tasks;

import android.app.ProgressDialog;
import android.content.Context;
import android.os.AsyncTask;

import com.google.gson.Gson;

import java.io.IOException;

import okhttp3.Call;
import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;
import xavier.ricardo.paulobet.Config;
import xavier.ricardo.paulobet.model.LoginRequest;
import xavier.ricardo.paulobet.MainActivity;
import xavier.ricardo.paulobet.model.User;

public class LoginTask extends AsyncTask<String, Void, Response> {
    private ProgressDialog progress;
    private MainActivity context;
    private String user;

    private static final MediaType JSON = MediaType.parse("application/json; charset=utf-8");

    public LoginTask(MainActivity context, String user, String password) {
        this.context = context;
        this.user = user;
    }

    @Override
    protected void onPreExecute() {
        progress = new ProgressDialog((Context) context);
        progress.setMessage("Aguarde...");
        progress.show();
        super.onPreExecute();
    }

    @Override
    protected void onPostExecute(Response response) {
        super.onPostExecute(response);
        progress.dismiss();
        try {
            context.onTaskResponse(response);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    protected void onProgressUpdate(Void... values) {
        super.onProgressUpdate(values);
    }

    @Override
    protected Response doInBackground(String... params) {
        try {
            Gson gson = new Gson();
            LoginRequest loginRequest = new LoginRequest();
            loginRequest.setLogin(user);
            String json = gson.toJson(loginRequest, LoginRequest.class);
            RequestBody body = RequestBody.create(JSON, json);
            Request request = new Request.Builder().url(Config.URL_LOGIN).post(body).build();
            OkHttpClient client = new OkHttpClient.Builder().build();
            Call call = client.newCall(request);
            return call.execute();
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }
}
