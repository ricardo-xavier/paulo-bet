package xavier.ricardo.paulobet.tasks;

import android.app.ProgressDialog;
import android.os.AsyncTask;

import com.google.gson.Gson;

import java.io.IOException;

import okhttp3.Call;
import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;
import xavier.ricardo.paulobet.ChangePasswordActivity;
import xavier.ricardo.paulobet.Config;
import xavier.ricardo.paulobet.model.ChangePasswordRequest;

public class ChangePasswordTask extends AsyncTask<String, Void, Response> {
    private ProgressDialog progress;
    private ChangePasswordActivity context;
    private String user;
    private String password;
    private String token;

    private static final MediaType JSON = MediaType.parse("application/json; charset=utf-8");

    public ChangePasswordTask(ChangePasswordActivity context, String user, String password, String token) {
        this.context = context;
        this.user = user;
        this.password = password;
        this.token = token;
    }

    @Override
    protected void onPreExecute() {
        progress = new ProgressDialog(context);
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
            ChangePasswordRequest changePasswordRequest = new ChangePasswordRequest();
            changePasswordRequest.setLogin(user);
            changePasswordRequest.setPassword(password);
            changePasswordRequest.setToken(token);
            String json = gson.toJson(changePasswordRequest, ChangePasswordRequest.class);
            RequestBody body = RequestBody.create(JSON, json);
            Request request = new Request.Builder().url(Config.URL_LOGIN).patch(body).build();
            OkHttpClient client = new OkHttpClient.Builder().build();
            Call call = client.newCall(request);
            return call.execute();
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }
}
