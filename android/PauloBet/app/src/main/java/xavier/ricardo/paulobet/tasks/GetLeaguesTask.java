package xavier.ricardo.paulobet.tasks;

import android.app.ProgressDialog;
import android.content.Context;
import android.os.AsyncTask;

import java.io.IOException;

import okhttp3.Call;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import xavier.ricardo.paulobet.Config;
import xavier.ricardo.paulobet.SelectLeagueActivity;

public class GetLeaguesTask extends AsyncTask<String, Void, Response> {
    private ProgressDialog progress;
    private SelectLeagueActivity context;
    private String user;

    public GetLeaguesTask(SelectLeagueActivity context, String user) {
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
            Request request = new Request.Builder().url(Config.URL_GET_LEAGUES + user).get().build();
            OkHttpClient client = new OkHttpClient.Builder().build();
            Call call = client.newCall(request);
            return call.execute();
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }
}
